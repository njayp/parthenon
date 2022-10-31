import logo from "./logo.svg";
import "./App.css";
import { BFCaller } from "./BFCaller";
import { Header } from "./Header";
import { RouteProps } from "react-router-dom";
import { Geo } from "./Geo";

// buttonLabel is used as key, so each should be unique
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
  routeProps: {
    path: "*",
    element: <img src={"/fof.png"} alt="404 dragon" width={"50%"} />,
  },
};

export const geoPage: Page = {
  buttonLabel: "Geolocation",
  routeProps: { path: "geo", element: <Geo /> },
};

export const homePage: Page = {
  buttonLabel: "Home",
  routeProps: { path: "/", element: <Header /> },
  subPages: [logoPage, fofPage, bfPage, geoPage],
};
