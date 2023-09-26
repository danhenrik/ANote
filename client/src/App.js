import "./App.css";
import MainContent from "./pages/MainContent";
import { ThemeProvider } from "@mui/material/styles";
import theme from "./theme";
import { useState } from "react";
import NavBar from "./components/header/NavBar";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import NotFound from "./pages/NotFound";
import CompleteSignupForm from "./components/AccessControl/Signup/CompleteSignupForm";

function App() {
  const [open, setOpen] = useState(false);

  return (
    <Router>
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
    </Router>
  );
}

export default App;
