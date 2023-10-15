import "./App.css";
import MainContent from "./pages/MainContent";
import { ThemeProvider } from "@mui/material/styles";
import { GoogleOAuthProvider } from "@react-oauth/google";
import defaultTheme from "./theme";
import { useState } from "react";
import NavBar from "./components/header/NavBar/NavBar";
import { BrowserRouter as Router } from "react-router-dom";
import config from "./configs.json";
import { AuthProvider } from "./store/auth-context";
import Routes from "./Routes/Routes";
import { ModalProvider } from "./store/modal-context";

function App() {
  const [open, setOpen] = useState(false);

  return (
    <Router>
      <AuthProvider>
        <GoogleOAuthProvider clientId={config.GOOGLE_CLIENT_ID}>
          <ThemeProvider theme={defaultTheme}>
            <ModalProvider>
              <NavBar open={open} setOpen={setOpen} />
              <MainContent open={open}>
                <Routes />
              </MainContent>
            </ModalProvider>
          </ThemeProvider>
        </GoogleOAuthProvider>
      </AuthProvider>
    </Router>
  );
}

export default App;
