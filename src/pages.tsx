import logo from "./logo.svg";
import "./App.css";
import { BFCaller } from "./BFCaller";
import { Header } from "./Header";
import { RouteProps } from "react-router-dom";

export type Page = {
  buttonLabel: string;
  routeProps: RouteProps;
  subPages?: Page[];
};

export const logoPage: Page = {
  buttonLabel: "Logo",
  routeProps: {
    index: true,
    element: <img src={logo} className="App-logo" alt="logo" />,
  },
};

export const bfPage: Page = {
  buttonLabel: "Boyfriend",
  routeProps: { path: "bfcaller", element: <BFCaller /> },
};

export const fofPage: Page = {
  buttonLabel: "404",
  routeProps: { path: "*", element: <>🤷 error 404 🤷</> },
};

/*
export const defaultRedirect: Page = {
  buttonLabel: "Error",
  routeProps: { path: "*", element: <Navigate to="404" /> },
};
*/

export const homePage: Page = {
  buttonLabel: "Home",
  routeProps: { path: "/", element: <Header /> },
  subPages: [logoPage, bfPage, fofPage],
};
