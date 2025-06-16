import { forwardRef } from 'react';
import { Icon } from '@iconify/react';
import { Box, type SxProps, type Theme } from '@mui/material';
import type { ComponentPropsWithoutRef, ComponentRef, ElementRef } from 'react';

type IconifyProps = {
  sx?: SxProps<Theme>;
  width?: number | string;
  icon: string;
};

const Iconify = forwardRef<ComponentPropsWithoutRef<typeof Box>, IconifyProps>(
  ({ icon, width = 20, sx, ...other }, ref) => {
    return (
      <Box
        ref={ref}
        sx={{ width, height: width, ...sx }}
        {...other}
      >
        <Icon icon={icon} width={width} height={width} />
      </Box>
    );
  }
);

export default Iconify;