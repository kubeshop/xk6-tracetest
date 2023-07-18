import { check } from "k6";
import { textSummary } from "https://jslib.k6.io/k6-summary/0.0.2/index.js";
import { Http, Tracetest } from "k6/x/tracetest";
import { sleep } from "k6";

export const options = {
  vus: 1,
  duration: "5s",
  thresholds: {
    http_req_duration: ["p(95)<1"], // 95% of requests should be below 200ms
  },
};
let pokemonId = 6; //charizard
const http = new Http();
const testId = "kc_MgKoVR";
const tracetest = Tracetest();

export default function () {
  const url = "http://localhost:8081/pokemon/import";
  const payload = JSON.stringify({
    id: pokemonId,
  });
  const params = {
    headers: {
      "Content-Type": "application/json",
    },
    tracetest: {
      testId,
    },
  };

  const response = http.post(url, payload, params);

  check(response, {
    "is status 200": (r) => r.status === 200,
    "body matches de id": (r) => JSON.parse(r.body).id === pokemonId,
  });

  pokemonId += 1;
  sleep(1);
}

// enable this to return a non-zero status code if a tracetest test fails
export function teardown() {
  tracetest.validateResult();
}

export function handleSummary(data) {
  // combine the default summary with the tracetest summary
  const tracetestSummary = tracetest.summary();
  const defaultSummary = textSummary(data);
  const summary = `
    ${defaultSummary}
    ${tracetestSummary}
  `;

  return {
    stderr: summary,
    "tracetest.json": tracetest.json(),
  };
}
