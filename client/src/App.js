import "./App.css";
import MainContent from "./pages/MainContent";
import { ThemeProvider } from "@mui/material/styles";
import { GoogleOAuthProvider } from "@react-oauth/google";
import theme from "./theme";
import { useState } from "react";
import NavBar from "./components/header/NavBar";
import { BrowserRouter as Router } from "react-router-dom";
import config from "./configs.json";
import { AuthProvider } from "./store/auth-context";
import Routes from "./Routes/Routes";
function App() {
  const [open, setOpen] = useState(false);

  return (
    <Router>
      <AuthProvider>
        <GoogleOAuthProvider clientId={config.GOOGLE_CLIENT_ID}>
          <ThemeProvider theme={theme}>
            <NavBar open={open} setOpen={setOpen} />
            <MainContent open={open}>
              <Routes />
            </MainContent>
          </ThemeProvider>
        </GoogleOAuthProvider>
      </AuthProvider>
    </Router>
  );
}

export default App;
