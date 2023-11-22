// cypress/integration/loginFormE2E.spec.js
describe("SignupForm E2E Test", () => {
  before(() => {
    cy.visit("/");
    cy.Login();
  });

  beforeEach(() => {
    cy.visit("/");
  });

  it("should open the create note modal and submit valid note", () => {
    cy.fixture("noteData.json").then((noteData) => {
      const { validNote } = noteData;
      cy.contains("Adicionar Nota").click();

      cy.get("#title").type(validNote.title);

      cy.get("#description").type(validNote.description);
      for (tag in validNote.tags) {
        cy.get("#tags").type(tag);
        cy.get("[data-testid='AddIcon']").click();
      }

      cy.get('[type="submit"]').click();
    });
  });

  // Add more test cases based on your component's behavior
});
