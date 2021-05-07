import React from "react";

type Props = {
  percentage: number
}

export default function Humidity(props: Props): JSX.Element {
  return <div className="humidity">
    <h3 className="value">{props.percentage}%</h3>
  </div>;
}
