import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
/** importing our pages */
import Tracks from './tracks';
import {ApolloProvider, ApolloClient, InMemoryCache} from "@apollo/client";

const client = new ApolloClient({
  uri: "http://localhost:4000",
  cache: new InMemoryCache(),
});

export default function Pages() {
  return (
    <ApolloProvider client={client}>
    <BrowserRouter>
      <Routes>
        <Route element={<Tracks />} path="/" />
      </Routes>
    </BrowserRouter>
    </ApolloProvider>
  );
}
