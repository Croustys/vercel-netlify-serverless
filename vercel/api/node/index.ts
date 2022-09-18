import { Pool } from "pg";
import dotenv from "dotenv";

dotenv.config();

export default function handler(req, res) {
  const connectionString = process.env.PG_URL;
  const pool = new Pool({
    connectionString,
  });
  const { name } = req.body;
  const query = {
    text: "INSERT INTO test_db (name) VALUES ($1)",
    values: [name],
  };
  pool.query(query, (err, r) => {
    if (err) {
      console.log(err.stack);
      return res.status(500).send("Unsuccesful database write");
    }
    pool.end();
    return res.status(200).send("Successful creation");
  });
}
