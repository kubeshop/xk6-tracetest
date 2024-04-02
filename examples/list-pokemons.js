import { Http, Tracetest } from "k6/x/tracetest";
import { sleep } from "k6";

export const options = {
  vus: 1,
  duration: "5s",
};

const http = new Http();
const tracetest = Tracetest();

export default function () {
  const url = "http://localhost:8081/pokemon?take=5";
  const response = http.get(url);

  tracetest.runTest(
    response.trace_id,
    {
      should_wait: true,
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
    stdout: tracetest.json(),
  };
}
