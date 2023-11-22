// cypress/integration/loginFormE2E.spec.js
describe("LoginForm E2E Test", () => {
  beforeEach(() => {
    cy.visit("/");
  });

  it("should open the login modal and submit valid credentials", () => {
    cy.fixture("loginData.json").then((loginData) => {
      const { validUser } = loginData;
      cy.contains("Login").click();

      cy.get("#login").type(validUser.login);
      cy.get("#password").type(validUser.password);

      cy.get('[type="submit"]').click();
      cy.contains("Feed de Notas").should("exist");
    });
  });

  it("should open the login modal and display an error for invalid credentials", () => {
    cy.fixture("loginData.json").then((loginData) => {
      const { invalidUser } = loginData;

      cy.contains("Login").click();

      cy.get("#login").type(invalidUser.login);
      cy.get("#password").type(invalidUser.password);

      cy.get('[type="submit"]').click();

      cy.contains("Falha no Login, dados incorretos").should("exist");
    });
  });

  // Add more test cases based on your component's behavior
});
