import { z } from 'zod';
import { ItemSchema } from './item';

export const CartItemSchema = z.object({
  id: z.number(),
  created_at: z.string().datetime().optional(),
  updated_at: z.string().datetime().optional(),
  cart_id: z.number(),
  item_id: z.number(),
  item: ItemSchema.optional(),
  quantity: z.number(),
});

export const CartDetailSchema = z.object({
  id: z.number(),
  created_at: z.string().datetime().optional(),
  updated_at: z.string().datetime().optional(),
  cart_id: z.number(),
  item_id: z.number(),
  item: ItemSchema.optional(),
  quantity: z.number(),
});

export type CartItem = z.infer<typeof CartItemSchema>;