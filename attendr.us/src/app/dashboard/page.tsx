/* eslint-disable react/jsx-no-undef */
import TrackedClasses from "@/components/TrackedClasses";
import AccountInformation from "@/components/AccountInformation";

const Dashboard = () => {
  return (
    <div className="m-3 mt-10">
      <div className="mb-6">
        <p className="text-3xl font-medium text-slate-900">Dashboard</p>
      </div>

      <div className="flex h-full flex-col items-center justify-center">
        <div className="mb-6 border-black">
          <p className="text-2xl font-medium text-slate-800">
            Account Information
          </p>
          {/* @ts-expect-error Async Server Component */}
          <AccountInformation />
        </div>

        <div className="w-full">
          <p className="text-2xl font-medium text-slate-800">Tracked Classes</p>
          {/* @ts-expect-error Async Server Component */}
          <TrackedClasses />
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
