"use client";

import FixedAmountForm from "@/components/forms/FixAmountForm";
import PercentageCategoryForm from "@/components/forms/PercentageCategoryForm";
import PercentageForm from "@/components/forms/PercentageForm";
import PointDiscountForm from "@/components/forms/PointDiscountForm";
import SeasonalForm from "@/components/forms/SeasonalForm";
import { useQueryCampaignOnDiscountCategory } from "@/query/campaign";
import { useQueryAvailableDiscountCategories } from "@/query/discount_category";
import { Campaign } from "@/types/campaign";
import { DiscountCategory } from "@/types/discount_category";
import React, { useState, useEffect } from "react";

const CampaignForm = () => {
  const [discountCategories, setDiscountCategories] = useState<DiscountCategory[]>([]);
  const [campaigns, setCampaigns] = useState<Campaign[]>([]);
  const [selectedDiscountCategory, setSelectedDiscountCategory] = useState<DiscountCategory | null>(null);
  const [selectedCampaign, setSelectedCampaign] = useState<Campaign | null>(null);

  const queryAvailableDiscountCategories = useQueryAvailableDiscountCategories();
  const queryCampaignOnDiscountType = useQueryCampaignOnDiscountCategory(selectedDiscountCategory?.id || 0);

  useEffect(() => {
    if (queryAvailableDiscountCategories.data) {
      setDiscountCategories(queryAvailableDiscountCategories.data);
      setCampaigns([])
      setSelectedCampaign(null)
    }
  }, [queryAvailableDiscountCategories.data]);

  useEffect(() => {
    if (queryCampaignOnDiscountType.data) {
      setCampaigns(queryCampaignOnDiscountType.data);
    } else {
      setCampaigns([]);
    }
  }, [queryCampaignOnDiscountType.data]);

  const renderForm = () => {
    if (!selectedCampaign) {
      return (
        <p className="text-gray-500 italic">
          Please select a campaign to display the form.
        </p>
      );
    }
    console.log(selectedDiscountCategory?.id!)

    switch (selectedCampaign?.name) {
      case "FIXED_AMOUNT":
        return <FixedAmountForm selectedDiscountCategoryId={selectedDiscountCategory?.id!}/>;
      case "PERCENTAGE":
        return <PercentageForm selectedDiscountCategoryId={selectedDiscountCategory?.id!}/>;
      case "PERCENTAGE_CATEGORY":
        return <PercentageCategoryForm selectedDiscountCategoryId={selectedDiscountCategory?.id!}/>;
      case "POINT":
        return <PointDiscountForm selectedDiscountCategoryId={selectedDiscountCategory?.id!}/>;
      case "SEASONAL":
        return <SeasonalForm selectedDiscountCategoryId={selectedDiscountCategory?.id!}/>;
      default:
        return (
          <p className="text-gray-500 italic">
            Please select a valid campaign to display the form.
          </p>
        );
    }
  };

  return (
    <div className="container mx-auto p-6 bg-white shadow-lg rounded-md">
      <h1 className="text-3xl font-bold mb-6 text-gray-800">Create Campaign on Discount Category</h1>

      <div className="mb-6">
        <label className="block text-lg font-medium text-gray-700 mb-2">Select Discount Category</label>
        <select
          onChange={(e) => {
            const selectedCategory = discountCategories.find(
              (category) => category.id.toString() === e.target.value
            );
            setSelectedDiscountCategory(selectedCategory || null);
            setSelectedCampaign(null); 
          }}
          value={selectedDiscountCategory?.id || ""}
          className="block w-full px-4 py-3 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
        >
          <option value="">
            -- Choose a Discount Category --
          </option>
          {discountCategories.map((category) => (
            <option key={category.id} value={category.id}>
              {category.name}
            </option>
          ))}
        </select>
        {!selectedDiscountCategory && (
          <p className="text-sm text-gray-500 mt-2">
            Choose a discount category to see related campaigns.
          </p>
        )}
      </div>

      <div className="mb-6">
        <label className="block text-lg font-medium text-gray-700 mb-2">Select Campaign</label>
        <select
          onChange={(e) => {
            const selectedCampaign = campaigns.find(
              (campaign) => campaign.id.toString() === e.target.value
            );
            setSelectedCampaign(selectedCampaign || null);
          }}
          value={selectedCampaign?.id || ""}
          className="block w-full px-4 py-3 border border-gray-300 rounded-md shadow-sm focus:ring-blue-500 focus:border-blue-500"
          disabled={!selectedDiscountCategory}
        >
          <option value="" disabled>
            {selectedDiscountCategory ? "-- Choose a Campaign --" : "Select a discount category first"}
          </option>
          {campaigns.map((campaign) => (
            <option key={campaign.id} value={campaign.id}>
              {campaign.name.replace("_", " ")}
            </option>
          ))}
        </select>
        {selectedDiscountCategory && campaigns.length === 0 && (
          <p className="text-sm text-gray-500 mt-2">No campaigns available for the selected category.</p>
        )}
      </div>

      {renderForm()}
    </div>
  );
};

export default CampaignForm;
