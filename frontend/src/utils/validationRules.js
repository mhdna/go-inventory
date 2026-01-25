export const itemValidationRules = {
  code: [(v) => !!v || "Code is required"],
  name: [(v) => !!v || "Name is required"],
  quantity: [
    (v) =>
      (v !== null && v !== undefined && v !== "") || "Quantity is required",
    (v) => v >= 0 || "Quantity cannot be negative",
  ],
};
