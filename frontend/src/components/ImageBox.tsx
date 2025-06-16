import { Box } from "@mui/material";
import React from "react";

type ImageBoxProps = {
  src: string;
  sx?: React.CSSProperties;
  alt?: string;
  width?: string | number;
  height?: string | number;
};

const ImageBox = ({ src, sx, alt, width, height }: ImageBoxProps) => {
  return (
    <Box
      component="img"
      src={src}
      sx={{
        width: width || "100%",
        height: height || "auto",
        ...sx,
      }}
      alt={alt || "Image"}
    />
  );
};

export default ImageBox;
