import React, { useState, useEffect } from "react";
import { usePlaidLink } from "react-plaid-link";
import "./App.scss";

function App() {
  const [token, setToken] = useState(null);
  const [accessToken, setAccessToken] = useState(null);
  const [error, setError] = useState(null);

  const onSuccess = async (publicToken) => {
    const response = await fetch("/api/exchange_public_token", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ public_token: publicToken }),
    });
    const { data, error } = await response.json();
    if (error) {
      return setError(error);
    }
    setAccessToken(data.access_token);
  };

  const createLinkToken = async () => {
    const response = await fetch("/api/create_link_token", {});
    const { data, error } = await response.json();
    if (error) {
      return setError(error);
    }
    return data.link_token;
  };

  const config = {
    token,
    onSuccess,
  };

  const { open, ready } = usePlaidLink(config);

  useEffect(() => {
    var ignore = false;
    if (token == null) {
      createLinkToken().then((linkToken) => {
        if (!ignore) {
          setToken(linkToken);
        }
      });
    }
    return () => {
      ignore = true;
    };
  }, [token]);

  return (
    <div>
      <button onClick={() => open()} disabled={!ready}>
        <strong>Link account</strong>
      </button>
      <div>Link Token: {token}</div>
      <div>Access Token: {accessToken}</div>
      <div>Error: {error}</div>
    </div>
  );
}

export default App;
