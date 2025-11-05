import { z } from "zod";

export const SignInSchema = z.object({
  email: z.email("Invalid email Address"),
  password: z.string().min(8, "Password must be at least 8 character"),
})

export const SignUpUserSchema = z.object({
  fullName: z.string().min(3, "Full Name must be at least 3 character"),
  email: z.email("Invalid email Address"),
  password: z.string().min(8, "Password must be at least 8 character"),
  confirmPassword: z.string().min(8, "Password must be at least 8 character"),
}).refine((data) => data.password == data.confirmPassword, {
  message: "password don't match",
  path: ['confirmPassword'],
})

export const SignUpResturantSchema = z.object({
  fullName: z.string().min(3, "Full Name must be at least 3 character"),
  email: z.email("Invalid email Address"),
  password: z.string().min(8, "Password must be at least 8 character"),
  resturantName: z.string().min(2, "Resturant Name must be at least 2 character"),
  confirmPassword: z.string().min(8, "Password must be at least 8 character"),
}).refine((data) => data.password == data.confirmPassword, {
  message: "password don't match",
  path: ['confirmPassword'],
})


export type SignInData = z.infer<typeof SignInSchema>;
export type SignUpUserData = z.infer<typeof SignUpUserSchema>;
export type SignUpResturantData = z.infer<typeof SignUpResturantSchema>;
