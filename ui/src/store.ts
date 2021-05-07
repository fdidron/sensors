import { flow, Instance, types as t } from "mobx-state-tree";
import { createContext, useContext } from "react";

const MOISTURE_SENSOR_FACTOR = 70;

function apiBase(): string {
  return import.meta.env.DEV === true
    ? "http://localhost:8080/"
    : window.location.href;
}
const Reading = t
  .model({
    value: t.string,
    createdAt: t.string,
  })
  .views((self) => {
    return {
      get lastUpdate(): string {
        let date = self.createdAt.slice(0, -1);
        date = date.replace("T", " @ ");
        return date;
      },
    };
  });

const Sensor = t
  .model({
    id: t.number,
    categoryId: t.number,
    name: t.string,
    description: t.string,
    unit: t.string,
    readings: t.array(Reading),
  })
  .views((self) => {
    return {
      get percentage(): number {
        let val = parseInt(self.readings[0].value, 10);
        if (self.unit === "Moisture") {
          val = ((val - 200) * 100) / MOISTURE_SENSOR_FACTOR;
        }
        return Math.round(val);
      },
    };
  });

const Store = t
  .model({
    sensors: t.array(Sensor),
  })
  .actions((self) => {
    const fetchSensors = flow(function* () {
      const res = yield fetch(`${apiBase()}sensors/`);
      const data = yield res.json();
      data.forEach((sensor: any) => {
        self.sensors.push(
          Sensor.create({
            id: sensor.id,
            categoryId: sensor.categoryId,
            name: sensor.name,
            description: sensor.description,
            unit: sensor.unit,
            readings: [
              Reading.create({
                value: sensor.lastReading.value,
                createdAt: sensor.lastReading.createdAt,
              }),
            ],
          })
        );
      });
    });

    const afterCreate = () => {
      fetchSensors();
    };

    return {
      afterCreate,
      fetchSensors,
    };
  })
  .views((self) => {
    return {
      get plantSensors() {
        return self.sensors.filter((sensor) => sensor.categoryId === 1);
      },
    };
  });

interface IStore extends Instance<typeof Store> {}

const store = Store.create();
const context = createContext(<IStore>{});
const useStore = () => useContext(context);

export { context, store, useStore };
