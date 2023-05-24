import { ReactNode } from "react";

const DashboardLayout = ({ children }: { children: ReactNode }) => {
  return (
    <section>
      <p>Account Information</p>
      <p>TODO: put stuff here like cool stats and notification preferences.</p>

      {children}
    </section>
  );
};

export default DashboardLayout;
