import { Http, Tracetest } from "k6/x/tracetest";
import { sleep } from "k6";

export const options = {
  vus: 1,
  duration: "5s",
};

const http = new Http();
const tracetest = Tracetest();
const definition = `type: Test
spec:
  id: k6-test
  name: K6 Test
  trigger:
    type: k6
  specs:
  - selector: span[tracetest.span.type="database" name="findMany pokeshop.pokemon" db.system="postgres" db.name="pokeshop" db.user="ashketchum" db.operation="findMany" db.sql.table="pokemon"]
    name: Correct db name
    assertions:
    - attr:db.name  =  "pokeshop"
`;

export default function () {
  const url = "http://localhost:8081/pokemon?take=5";
  const response = http.get(url);

  tracetest.runTest(
    response.trace_id,
    {
      should_wait: true,
      definition,
    },
    {
      url,
      method: "GET",
    }
  );

  sleep(1);
}

export function handleSummary() {
  return {
    stdout: tracetest.summary(),
  };
}
