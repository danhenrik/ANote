// cypress/integration/loginFormE2E.spec.js
describe("LoginForm E2E Test", () => {
  beforeEach(() => {
    // Visit the main page before each test
    cy.visit("/");
  });

  it("should open the login modal and submit valid credentials", () => {
    // Load valid user data from the fixture
    cy.fixture("loginData.json").then((loginData) => {
      const { validUser } = loginData;
      cy.contains("Login").click();
      cy.contains("Faça Login").should("exist");

      cy.get("#login").type(validUser.login);
      cy.get("#password").type(validUser.password);

      cy.get('[type="submit"]').click();
      cy.url().should("include", "/");
    });
  });

  it("should open the login modal and display an error for invalid credentials", () => {
    // Load invalid user data from the fixture
    cy.fixture("loginData.json").then((loginData) => {
      const { invalidUser } = loginData;

      // Open the login modal
      cy.contains("Login").click();
      cy.contains("Faça Login").should("exist");

      // Fill in the form with invalid credentials
      cy.get("#login").type(invalidUser.login);
      cy.get("#password").type(invalidUser.password);

      // Submit the form
      cy.get('[type="submit"]').click();

      // Add assertions based on the expected behavior after submitting the form
      // For example, check if an error message is displayed
      cy.contains("Falha no Login, dados incorretos").should("exist");

      // Add more assertions based on your component's behavior
    });
  });

  // Add more test cases based on your component's behavior
});
