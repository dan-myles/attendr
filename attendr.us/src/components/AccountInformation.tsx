import { currentUser } from "@clerk/nextjs";
import Image from "next/image";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

const AccountInformation = async () => {
  const user = await currentUser();
  if (!user) return null;

  console.log(user);

  return (
    <div className="flex flex-wrap xl:space-x-3">
      <div className="rounded-2xl transition ease-in-out hover:-translate-y-1 hover:shadow-2xl">
        <Card>
          <CardHeader>
            <div>
              <Image
                src={user.profileImageUrl}
                alt="Profile Image"
                width={64}
                height={64}
                className="float-left mr-8 rounded-full"
              />
              <p className="text-[110%] font-[400]">
                {user.firstName + " " + user.lastName}
              </p>
              <p className="text-[75%] font-[300] text-slate-700">{user.id}</p>
              <p className="text-[75%] font-[300] text-slate-700">
                {user.primaryPhoneNumberId}
              </p>
            </div>
            <CardTitle></CardTitle>
            <CardDescription></CardDescription>
          </CardHeader>
          <CardContent>
            <p className="-mt-5 mb-3 text-[115%] font-semibold">
              Notification Settings
            </p>
            <p className="text-[105%] text-gray-900">Email:</p>
            {user.emailAddresses.map((e, i) => (
              <div key={i}>
                <p className="text-gray-600">{e.emailAddress}</p>
              </div>
            ))}

            <p className="mt-4 text-[105%] text-gray-900">Phone:</p>
            {user.phoneNumbers.map((p, i) => (
              <div key={i}>
                <p className="text-gray-600">{p.phoneNumber}</p>
              </div>
            ))}

            <p className="mt-4 text-sm font-light">
              We will use your contact information to send you updates when a
              class is available!
            </p>
          </CardContent>
          <CardFooter></CardFooter>
        </Card>
      </div>
      <div className="rounded-2xl transition ease-in-out hover:-translate-y-1 hover:shadow-2xl">
        <Card>
          <CardHeader>
            <div>
              <Image
                src={user.profileImageUrl}
                alt="Profile Image"
                width={64}
                height={64}
                className="float-left mr-8 rounded-full"
              />
              <p className="text-[110%] font-[400]">
                {user.firstName + " " + user.lastName}
              </p>
              <p className="text-[75%] font-[300] text-slate-700">{user.id}</p>
              <p className="text-[75%] font-[300] text-slate-700">
                {user.primaryPhoneNumberId}
              </p>
            </div>
            <CardTitle></CardTitle>
            <CardDescription></CardDescription>
          </CardHeader>
          <CardContent>
            <p className="-mt-5 mb-3 text-[115%] font-semibold">
              Notification Settings
            </p>
            <p className="text-[105%] text-gray-900">Email:</p>
            {user.emailAddresses.map((e, i) => (
              <div key={i}>
                <p className="text-gray-600">{e.emailAddress}</p>
              </div>
            ))}

            <p className="mt-4 text-[105%] text-gray-900">Phone:</p>
            {user.phoneNumbers.map((p, i) => (
              <div key={i}>
                <p className="text-gray-600">{p.phoneNumber}</p>
              </div>
            ))}

            <p className="mt-4 text-sm font-light">
              We will use your contact information to send you updates when a
              class is available!
            </p>
          </CardContent>
          <CardFooter></CardFooter>
        </Card>
      </div>
    </div>
  );
};

export default AccountInformation;
