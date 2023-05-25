"use client";

import { FC } from "react";
import { SignInButton, SignUpButton, UserButton, useAuth } from "@clerk/nextjs";
import { usePathname } from "next/navigation";
import styles from "./TopButtons.module.css";

const TopButtons: FC = () => {
  const { isSignedIn } = useAuth();
  const pathname = usePathname();

  return (
    <div className={styles.userButtons}>
      {!isSignedIn && (
        <div>
          <SignUpButton mode="modal" redirectUrl="/dashboard">
            <button className={styles.topButton}>Sign Up</button>
          </SignUpButton>
          <SignInButton mode="modal" redirectUrl="/dashboard">
            <button className={styles.topButton}>Sign In</button>
          </SignInButton>
        </div>
      )}

      {isSignedIn && (
        <div className={styles.dashboardButtons}>
          {pathname === "/dashboard" && (
            <div>
              <a href="/">
                <button className={styles.topButton}>Home</button>
              </a>
            </div>
          )}

          {pathname !== "/dashboard" && (
            <div>
              <a href="/dashboard">
                <button className={styles.topButton}>Dashboard</button>
              </a>
            </div>
          )}
          <div className={styles.userButton}>
            <div
              className="
              transition ease-in-out hover:-translate-y-1 hover:shadow-2xl
              active:translate-y-0 active:shadow-none
              "
            >
              <UserButton
                appearance={{
                  elements: {
                    userButtonAvatarBox: {
                      width: 56,
                      height: 56,
                    },
                  },
                }}
              />
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default TopButtons;
