import React from "react";
import { Navigate, useRoutes } from "react-router-dom";
import CompleteSignupForm from "../components/AccessControl/Signup/CompleteSignupForm";
import PasswordRecovery from "../components/AccessControl/PasswordRecovery/PasswordRecovery";
import NotFound from "../pages/NotFound";
import VerificationCode from "../components/AccessControl/PasswordRecovery/VerficationCode";
import CompleteRecovery from "../components/AccessControl/PasswordRecovery/CompleteRecovery";
import Timeline from "../components/Timeline/NotesTimeline/Timeline";
import Communities from "../components/Timeline/Communities/Communities";

function Routes() {
  return useRoutes([
    {
      path: "/",
      element: <Navigate to='/timeline' />,
    },
    {
      path: "/timeline",
      element: <Timeline />,
    },
    {
      path: "/communities",
      element: <Communities />,
    },
    {
      path: "/community/:id",
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
