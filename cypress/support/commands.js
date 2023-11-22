// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --
// Cypress.Commands.add('login', (email, password) => { ... })
//
//
// -- This is a child command --
// Cypress.Commands.add('drag', { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add('dismiss', { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This will overwrite an existing command --
// Cypress.Commands.overwrite('visit', (originalFn, url, options) => { ... })
Cypress.Commands.add("uploadFile", (fileName, selector) => {
  cy.fixture(fileName).then((fileContent) => {
    cy.get(selector).upload({
      fileContent,
      fileName,
      mimeType: "application/octet-stream",
    });
  });
});
Cypress.Commands.add("Login", () => {
  cy.fixture("loginData.json").then((loginData) => {
    const { validUser } = loginData;
    cy.contains("Login").click();

    cy.get("#login").type(validUser.login);
    cy.get("#password").type(validUser.password);

    cy.get('[type="submit"]').click();
  });
});

Cypress.Commands.add("AdicionaNota", () => {
  cy.fixture("notaData.json").then((notaData) => {
    const { validUser } = loginData;
    cy.contains("Login").click();

    cy.get("#login").type(validUser.login);
    cy.get("#password").type(validUser.password);

    cy.get('[type="submit"]').click();
  });
});
