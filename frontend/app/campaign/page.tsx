"use client";

import DiscountCard from "@/components/cards/DiscountCard";
import CampaignForm from "@/components/forms/CampaignForm";
import { useMutationDeleteDiscount, useQueryDiscounts } from "@/query/discount";
import React from "react";

const CampaignPage = () => {
  const { data: discounts } = useQueryDiscounts();
  const mutationDeleteDiscount = useMutationDeleteDiscount()
  return (
    <div className="bg-white min-h-screen">
      {/* Form Component */}
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 mt-10">
        <CampaignForm />
      </div>

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 mt-10">
        <h3 className="text-3xl font-semibold text-gray-800 text-center mb-8">
          Active Discount Campaigns
        </h3>

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
          {discounts?.map((discount) => (
            <DiscountCard key={discount.id} discount={discount} onRemove={async() => await mutationDeleteDiscount.mutateAsync(discount.id)}/>
          ))}
        </div>
      </div>
    </div>
  );
};

export default CampaignPage;
