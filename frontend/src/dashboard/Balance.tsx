import { Box, Button, Stack, Typography } from "@mui/material";

const Balance = () => {
  return (
    <Box>
      <Stack direction="row" justifyContent="center" alignItems="center">
        <Box
          sx={{
            py: 0.35,
            bgcolor: "primary.main",
            borderTopLeftRadius: 4,
            borderBottomLeftRadius: 4,
          }}
        >
          <Typography variant="overline" sx={{ p: 1 }}>
            $10,000
          </Typography>
        </Box>
        <Button
          variant="contained"
          color="info"
          size="small"
          sx={{ borderTopLeftRadius: 0, borderBottomLeftRadius: 0 }}
        >
          Wallet
        </Button>
      </Stack>
    </Box>
  );
};

export default Balance;
