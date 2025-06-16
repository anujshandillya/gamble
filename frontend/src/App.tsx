import { Route, BrowserRouter as Router, Routes } from "react-router-dom";
import DashboardLayout from "./dashboard/DashboardLayout";
import { createTheme, CssBaseline, ThemeProvider } from "@mui/material";
import { themeSettings } from "./theme";
import Home from "./zero/home/Home";

function App() {
  const theme = createTheme(themeSettings);
  return (
    <Router>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <DashboardLayout>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/casino/home" element={<Home />} />
            <Route path="/casino/games/dice" element={<Home />} />
            <Route path="/casino/games/limbo" element={<Home />} />
            <Route path="/casino/games/wheel" element={<Home />} />
            <Route path="/casino/games/dragontower" element={<Home />} />
            <Route path="/casino/games/cointoss" element={<Home />} />
            <Route path="/casino/games/mines" element={<Home />} />
          </Routes>
        </DashboardLayout>
      </ThemeProvider>
    </Router>
  );
}

export default App;
