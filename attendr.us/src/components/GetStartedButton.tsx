"use client";

import { SignUpButton, useAuth } from "@clerk/nextjs";
import styles from "./GetStartedButton.module.css";
import { FC } from "react";

const GetStartedButton: FC = () => {
  const { isSignedIn } = useAuth();

  return (
    <div>
      {!isSignedIn && (
        <div className={styles.getStartedSection}>
          <SignUpButton mode="modal">
            <button className={styles.getStartedButton}>try it out free</button>
          </SignUpButton>
        </div>
      )}

      {isSignedIn && (
        <div className={styles.getStartedSection}>
          <a href="/dashboard">
            <button className={styles.getStartedButton}>
              go to the dashboard
            </button>
          </a>
        </div>
      )}
    </div>
  );
};

export default GetStartedButton;
