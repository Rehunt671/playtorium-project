import { z } from "zod";
import { ItemCategorySchema } from "./item";

export const PercentageCategoryDiscountSchema = z.object({
  id: z.number().int(),
  item_category_id: z.number().int(),
  item_category: ItemCategorySchema,
  percentage: z.number(),
});

export const CreatePercentageCategoryDiscountSchema = z.object({
  item_category_id: z.number().int(),
  percentage: z.number(),
});

export type PercentageCategoryDiscount = z.infer<typeof PercentageCategoryDiscountSchema>;
