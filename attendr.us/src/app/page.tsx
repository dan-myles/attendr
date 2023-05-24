import styles from "./page.module.css";
import { FC } from "react";
import GetStartedButton from "@/components/GetStartedButton";

const Home: FC = () => {
  return (
    <main className={styles.main}>
      <div className={styles.mainSection}>
        <p className={styles.mainSlogan}>never miss a class again 🚀</p>
        <p className={styles.mainUnderSlogan}>
          let attendr keep track of full classes, we&apos;ll send you a text
          when a spot opens up!
        </p>
      </div>

      <GetStartedButton />

      <div className={styles.footer}>
        <p>© 2023 attendr | all rights reserved</p>
      </div>
    </main>
  );
};

export default Home;
