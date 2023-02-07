require("dotenv").config();
const express = require("express");
const { Configuration, PlaidApi, PlaidEnvironments } = require("plaid");

const PRODUCTS = [
  "assets",
  // "auth",
  // "employment",
  "identity",
  // "income_verification",
  // "identity_verification",
  "investments",
  "liabilities",
  // "payment_initiation",
  // "standing_orders",
  "transactions",
  // "transfer",
];

const COUNTRY_CODES = [
  "US",
  "GB",
  "ES",
  "NL",
  "FR",
  "IE",
  "CA",
  "DE",
  "IT",
  "PL",
  "DK",
  "NO",
  "SE",
  "EE",
  "LT",
  "LV",
];

const createApp = () => {
  if (!process.env.PLAID_CLIENT_ID) {
    throw new Error("Missing PLAID_CLIENT_ID");
  }
  if (!process.env.PLAID_SECRET) {
    throw new Error("Missing PLAID_SECRET");
  }
  const app = express();
  app.use(express.urlencoded({ extended: false }));
  app.use(express.json());

  const config = new Configuration({
    basePath: PlaidEnvironments[process.env.PLAID_ENV || "sandbox"],
    baseOptions: {
      headers: {
        "PLAID-CLIENT-ID": process.env.PLAID_CLIENT_ID,
        "PLAID-SECRET": process.env.PLAID_SECRET,
      },
    },
  });

  const client = new PlaidApi(config);
  app.get("/api/create_link_token", async (req, res, next) => {
    try {
      const tokenResponse = await client.linkTokenCreate({
        user: { client_user_id: "cq-source-plaid" },
        client_name: "CloudQuery Source Plaid",
        language: "en",
        products: PRODUCTS,
        country_codes: COUNTRY_CODES,
      });
      res.json({ data: tokenResponse.data });
    } catch (e) {
      console.log(e.response.data);
      res.status(500).json({ error: "Failed to create link token" });
      next();
    }
  });

  app.post("/api/exchange_public_token", async (req, res, next) => {
    try {
      const exchangeResponse = await client.itemPublicTokenExchange({
        public_token: req.body.public_token,
      });
      res.json({ data: { access_token: exchangeResponse.data.access_token } });
    } catch (e) {
      console.log(e.response.data);
      res.status(500).json({ error: "Failed to exchange token" });
      next();
    }
  });

  return app;
};

const app = createApp();
const port = 8080;
app.listen(port).on("listening", () => {
  console.log("Listening on port " + port + "...");
});
