import { z } from 'zod';

export const ItemSchema = z.object({
  id: z.number(),
  name: z.string(),
  price: z.number(),
  image_name: z.string().optional(),
  description: z.string().optional(),
  item_category_id: z.number().optional(),
  item_category: z.object({
    id: z.number(),
    name: z.string(),
  }).optional(),
  stock: z.number().optional(),
});

export const ItemCategorySchema = z.object({
  id: z.number().int(),
  name: z.string(),
});


export type Item = z.infer<typeof ItemSchema>;
export type ItemCategory = z.infer<typeof ItemCategorySchema>;
