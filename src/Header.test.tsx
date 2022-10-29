import React from "react";
import { render, screen } from "@testing-library/react";
import { Header } from "./Header";
import { MemoryRouter } from "react-router-dom";

describe("header tests", () => {
  render(
    <MemoryRouter>
      <Header />
    </MemoryRouter>
  );
  it("has a home button", () => {
    const linkElement = screen.getByText(/Home/i);
    expect(linkElement).toBeInTheDocument();
  });
});
