import { z } from 'zod';

export const LoginCredentialsSchema = z.object({
  username: z.string(),
  password: z.string(),
});

export type LoginCredentials = z.infer<typeof LoginCredentialsSchema>;

export const RegisterCredentialsSchema = z.object({
  name: z.string(),
  username: z.string(),
  password: z.string(),
});

export type RegisterCredentials = z.infer<typeof RegisterCredentialsSchema>;