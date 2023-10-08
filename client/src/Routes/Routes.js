import React from "react";
import { useRoutes } from "react-router-dom";
import CompleteSignupForm from "../components/AccessControl/Signup/CompleteSignupForm";
import PasswordRecovery from "../components/AccessControl/PasswordRecovery/PasswordRecovery";
import NotFound from "../pages/NotFound";
import VerificationCode from "../components/AccessControl/PasswordRecovery/VerficationCode";
import CompleteRecovery from "../components/AccessControl/PasswordRecovery/CompleteRecovery";
import Timeline from "../components/NotesTimeline/Timeline";

function Routes() {
  return useRoutes([
    {
      path: "/",
      element: <Timeline />,
    },
    {
      path: "/signup",
      element: <CompleteSignupForm />,
    },
    {
      path: "/passwordrecovery",
      element: <PasswordRecovery />,
    },
    {
      path: "/passwordrecovery/verificationcode",
      element: <VerificationCode />,
    },
    {
      path: "/passwordrecovery/changepassword",
      element: <VerificationCode />,
    },
    {
      path: "/passwordrecovery/completerecovery",
      element: <CompleteRecovery />,
    },
    {
      path: "*",
      element: <NotFound />,
    },
  ]);
}

export default Routes;
