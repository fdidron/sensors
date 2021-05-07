import React from "react";
import { observer } from "mobx-react-lite";

import { useStore } from "./store";
import Humidity from "./humidity";

function App(): JSX.Element {
  const store = useStore();
  return (
    <div className="frame">
      <h1>Plant sensors</h1>
      <div className="sensors">
        {store.plantSensors.map((sensor) => {
          let sensorClass = "sensor";
          if (sensor.unit === "Moisture") sensorClass += " drop";
          return (
            <div className={sensorClass}>
              <div className="body">
                <Humidity percentage={sensor.percentage} />
                <div className="last-update">
                  Updated at: {sensor.readings[0].lastUpdate}
                </div>
                <h2 className="description">{sensor.description}</h2>
              </div>
            </div>
          );
        })}
      </div>
      <h1>Home sensors</h1>
    </div>
  );
}

export default observer(App);
