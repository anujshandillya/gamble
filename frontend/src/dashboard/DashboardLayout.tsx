import type React from "react";
import Header from "../components/Header";
import Footer from "../components/Footer";
import Main from "../components/Main";

type LayoutProps = {
  children: React.ReactNode;
};

const DashboardLayout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <>
      <Header />
      <Main>{children}</Main>
      <Footer />
    </>
  );
};

export default DashboardLayout;
