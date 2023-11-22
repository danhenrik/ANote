// cypress/integration/loginFormE2E.spec.js
describe("SignupForm E2E Test", () => {
  beforeEach(() => {
    cy.visit("/");
  });

  it("should open the signup modal and submit valid email", () => {
    cy.fixture("signupData.json").then((signupData) => {
      const { validUser } = signupData;
      cy.contains("Cadastre-se").click();

      cy.get("#email").type(validUser.email);

      cy.get('[type="submit"]').click();

      cy.get("#username").type(validUser.login);
      cy.get("#password").type(validUser.password);
      cy.get("#confirmPassword").type(validUser.password);
      cy.get("input[type=file]").selectFile("./cypress/imgs/your_image.jpg", {
        force: true,
      });
      cy.get('[type="submit"]').click();
      cy.contains("Feed de Notas").should("exist");
    });
  });

  // Add more test cases based on your component's behavior
});
