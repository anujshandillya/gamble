import { Box, Stack, Typography, Link, IconButton } from "@mui/material";
import ImageBox from "./ImageBox";
import Iconify from "./iconify";

const Footer = () => {
  return (
    <Box
      sx={{
        width: "100%",
        boxShadow: 1,
        bgcolor: (theme) => theme.palette.background.paper,
        padding: 2,
      }}
    >
      <Stack
        direction="row"
        justifyContent="space-evenly"
        alignItems="center"
      >
        {/* Logo and Branding */}
        <Stack
          direction="row"
          alignItems="center"
          gap={1}
          sx={{ height: "100%" }}
        >
          <ImageBox src="/logo-nobg.png" width={64} height={64} sx={{}} />
          <Typography
            variant="subtitle1"
            color="primary.main"
            sx={{ display: { xs: "none", sm: "block" } }}
          >
            AceUp
          </Typography>
        </Stack>

        {/* Navigation Links */}
        <Stack
          direction="row"
          justifyContent="flex-end"
          alignItems="center"
          gap={{xs: 1, sm: 2}}
          sx={{ height: "100%", paddingRight: "16px" }}
        >
          <Link href="/about" color="primary" underline="hover">
            About Us
          </Link>
          <Link href="/contact" color="primary" underline="hover">
            Contact
          </Link>
          <Link href="/privacy" color="primary" underline="hover">
            Privacy Policy
          </Link>
          <Link href="/terms" color="primary" underline="hover">
            Terms of Service
          </Link>
        </Stack>

        {/* Social Media Icons */}
        <Stack direction="row" gap={1}>
          <IconButton
            href="https://facebook.com"
            target="_blank"
            rel="noopener"
            color="primary"
          >
            <Iconify icon="ic:baseline-facebook" />
          </IconButton>
          <IconButton
            href="https://twitter.com"
            target="_blank"
            rel="noopener"
            color="primary"
          >
            <Iconify icon="pajamas:twitter" />
          </IconButton>
          <IconButton
            href="https://instagram.com"
            target="_blank"
            rel="noopener"
            color="primary"
          >
            <Iconify icon="mdi:instagram" />
          </IconButton>
        </Stack>
      </Stack>

      {/* Copyright Section */}
      <Box sx={{ textAlign: "center", marginTop: 2 }}>
        <Typography variant="body2" color="primary.dark">
          © {new Date().getFullYear()} AceUp. All rights reserved.
        </Typography>
      </Box>
    </Box>
  );
};

export default Footer;