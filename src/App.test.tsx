import React from "react";
import { render, screen } from "@testing-library/react";
import { Header } from "./Header";
import { MemoryRouter } from "react-router-dom";

describe("render header", () => {
  render(
    <MemoryRouter>
      <Header />
    </MemoryRouter>
  );
  it("home button exists", () => {
    const linkElement = screen.getByText(/Home/i);
    expect(linkElement).toBeInTheDocument();
  });
});
