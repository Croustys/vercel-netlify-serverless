import { Handler } from "@netlify/functions";

const handler: Handler = async (event, context) => {
  return {
    statusCode: 200,
    body: "Netlify Node Serveless Performance test",
  };
};

export { handler };