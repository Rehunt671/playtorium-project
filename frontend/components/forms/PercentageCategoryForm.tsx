import { useState, useEffect } from "react";
import { CreateDiscount } from "@/types/discount";
import { useMutationCreateDiscount } from "@/query/discount";
import { useQueryItemCategories } from "@/query/item_category";

interface PercentageCategoryFormProps {
  selectedDiscountCategoryId: number;
}

const PercentageCategoryForm = ({ selectedDiscountCategoryId }: PercentageCategoryFormProps) => {
  const [itemCategoryId, setItemCategoryId] = useState<number>(0);
  const [percentage, setPercentage] = useState<number>(0);
  const [error, setError] = useState('');
  const { mutateAsync } = useMutationCreateDiscount();
  const { data: itemCategories,isPending } = useQueryItemCategories();

  const onSubmitDiscount = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (itemCategoryId === 0) {
      setError("Please select an item category.");
      return; 
    }
    
    if (percentage === 0) {
      setError("Percentage must > 0");
      return; 
    }

    setError("");

    const discountData: CreateDiscount = {
      percentage_category: {
        item_category_id: itemCategoryId,
        percentage: percentage,
      },
      discount_category_id: selectedDiscountCategoryId,
    };

    try {
      await mutateAsync(discountData);
    } catch (err) {
      console.error("Error creating discount", err);
    }
  };

  return (
    <form onSubmit={onSubmitDiscount}>
      <h2 className="text-xl font-bold text-gray-800 mb-4">Percentage Category Discount</h2>

      {/* Item Category Dropdown */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-2">
          Item Category
        </label>
        <select
          name="itemCategoryId"
          value={itemCategoryId}
          onChange={(e) => setItemCategoryId(Number(e.target.value))}
          className="block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
          required
        >
          <option value={0} disabled>
            Select a category
          </option>
          {itemCategories?.map((category) => (
            <option key={category.id} value={category.id}>
              {category.name}
            </option>
          ))}
        </select>
        {/* Display error message if no category is selected */}
        {error && <p className="text-red-500 text-sm mt-2">{error}</p>}
      </div>

      {/* Percentage Input */}
      <div className="mb-4">
        <label className="block text-sm font-medium text-gray-700 mb-2">
          Percentage
        </label>
        <input
          type="number"
          name="percentage"
          value={percentage}
          onChange={(e) => setPercentage(Number(e.target.value))}
          placeholder="Enter percentage discount"
          className="block w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
          required
        />
      </div>

      {error && <p className="text-red-500 text-sm mb-4">{error}</p>}

      <button
        type="submit"
        className="w-full bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        disabled={isPending}
      >
        {isPending ? "Creating..." : "Apply Discount"}
      </button>
    </form>
  );
};

export default PercentageCategoryForm;
