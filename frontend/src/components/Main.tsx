import { Box } from "@mui/material";
import React from "react";

type MainProps = {
  children: React.ReactNode;
};

const Main: React.FC<MainProps> = ({ children }) => {
  return (
    <Box
      sx={{
        pt: "60px",
      }}
    >
      {children}
    </Box>
  );
};

export default Main;
