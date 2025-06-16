import { Box, Button, IconButton, Stack, Typography } from "@mui/material";
import ImageBox from "./ImageBox";
import Iconify from "./iconify";
import Balance from "../dashboard/Balance";

const Header = () => {
  return (
    <Box
      sx={{
        width: "100%",
        height: "60px",
        boxShadow: 1,
        position: "fixed",
        top: 0,
        bgcolor: (theme) => theme.palette.background.paper,
      }}
    >
      <Stack
        direction="row"
        justifyContent="space-evenly"
        alignItems="center"
        sx={{ height: "100%" }}
      >
        <Stack
          direction="row"
          alignItems="center"
          gap={1}
          sx={{ height: "100%" }}
        >
          <ImageBox src="/logo-nobg.png" width={64} height={64} sx={{}} />
          <Typography variant="subtitle1" color="primary.main" sx={{ display: { xs: "none", sm: "block" } }}>AceUp</Typography>
        </Stack>
        <Balance />
        <Stack
          direction="row"
          justifyContent="flex-end"
          alignItems="center"
          gap={{xs: 1, sm: 2}}
          sx={{ height: "100%", paddingRight: "16px" }}
        >
            <Button variant="text" color="primary" startIcon={<Iconify icon="iconamoon:search" />}>
                Search
            </Button>
            <IconButton color="primary">
                <Iconify icon="iconamoon:profile-fill" />
            </IconButton>
            <IconButton color="primary">
                <Iconify icon="material-symbols:notifications" />
            </IconButton>
            <IconButton color="primary">
                <Iconify icon="material-symbols:receipt" />
            </IconButton>
        </Stack>
      </Stack>
    </Box>
  );
};

export default Header;
