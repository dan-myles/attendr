import styles from "./page.module.css";
import { FC } from "react";
import GetStartedButton from "@/components/GetStartedButton";

const Home: FC = () => {
  return (
    <main>
      <div className="flex flex-col items-center justify-center">
        <div className="mt-56">
          <p className="text-6xl">never miss a class again 🚀</p>
          <p className="mt-2 text-xl">
            let attendr keep track of full classes, we&apos;ll send you a text
            when a spot opens up!
          </p>
        </div>

        <div className="mb-48 mt-10">
          <GetStartedButton />
        </div>

        <div></div>

        <div className="text-l fixed bottom-0 w-full text-center">
          <p>© 2023 attendr | all rights reserved</p>
        </div>
      </div>
    </main>
  );
};

export default Home;
