import React from "react";
import { TooltipProvider } from "../ui/tooltip";

type ProvidersProps = React.PropsWithChildren<{}>;

const Providers: React.FC<ProvidersProps> = ({ children }) => {
  return <TooltipProvider>{children}</TooltipProvider>;
};

export default Providers;
