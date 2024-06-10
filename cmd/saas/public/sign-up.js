window.addEventListener("load", function () {
  const { createApp } = Vue;
  createApp({
    data() {
      return {
        email: "",
        password: "",
        error: "",
      };
    },
    watch: {
      password() {
        this.error = "";
      },
      email() {
        this.error = "";
      },
    },
    methods: {
      async signUp() {
        if (this.email === "" || this.password === "") {
          this.error = "Email and password are required";
          return;
        }

        const data = {
          email: this.email,
          emailVisibility: true,
          password: this.password,
          passwordConfirm: this.password,
        };

        const response = await fetch("/api/collections/users/records", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        });

        if (!response.ok) {
          const respJson = await response.json();
          this.error = respJson.message;
          return;
        }

        const requestVerification = await fetch(
          "/api/collections/users/request-verification",
          {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({ email: this.email }),
          }
        );

        if (requestVerification.ok) {
          window.location.href = "/auth/sign-in";
        } else {
          const requestVerificationJson = await response.json();
          this.error = requestVerificationJson.message;
        }
      },
    },
  }).mount("#sign-up-form");
});
