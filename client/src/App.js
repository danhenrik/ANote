import "./App.css";
import MainContent from "./pages/MainContent";
import { ThemeProvider } from "@mui/material/styles";
import { GoogleOAuthProvider } from "@react-oauth/google";
import theme from "./theme";
import { useState } from "react";
import NavBar from "./components/header/NavBar";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import NotFound from "./pages/NotFound";
import CompleteSignupForm from "./components/AccessControl/Signup/CompleteSignupForm";
import config from "./configs.json";
import { AuthProvider } from "./store/auth-context";

function App() {
  const [open, setOpen] = useState(false);

  return (
    <Router>
      <AuthProvider>
        <GoogleOAuthProvider clientId={config.GOOGLE_CLIENT_ID}>
          <ThemeProvider theme={theme}>
            <NavBar open={open} setOpen={setOpen} />
            <MainContent open={open}>
              <Routes>
                <Route path='/' />
                <Route path='/signup' element={<CompleteSignupForm />} />
                <Route path='*' element={<NotFound />} />
              </Routes>
            </MainContent>
          </ThemeProvider>
        </GoogleOAuthProvider>
      </AuthProvider>
    </Router>
  );
}

export default App;
