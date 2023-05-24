import TrackedClasses from "@/components/TrackedClasses";
import { FC } from "react";

const Dashboard: FC = () => {
  return (
    <div>
      <h3>Dashboard</h3>
      {/* @ts-expect-error Async Server Component */}
      <TrackedClasses />
    </div>
  );
};

export default Dashboard;
