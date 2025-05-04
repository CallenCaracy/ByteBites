import { useState } from "react";
import { useMutation } from "@apollo/client";
import { ADD_MENU_ITEM } from "../graphql/Menuqueries"; // adjust the path if needed
import { supabase } from "../utils/supabaseClient";
import { useNavigate, useParams } from "react-router-dom";
import Navbar from "../components/NavBar";

const AddMenu = () => {
    const navigate = useNavigate();
    const { userId } = useParams();

    const [formData, setFormData] = useState({
      name: "",
      description: "",
      price: "",
      category: "",
      discount: "",
      availability_status: true,
      image_url: "https://hzjjmfwrtvqjwxunfcue.supabase.co/storage/v1/object/public/pictures/menupic/placeholder.jpg",
    });

    const [fileUploading, setFileUploading] = useState(false);

    const [addMenuItem, { loading, error }] = useMutation(ADD_MENU_ITEM, {
        onCompleted: (data) => {
            const newMenuId = data?.createMenuItem?.id;
            console.log("Created:", data?.createMenuItem?.id);

            setFormData({
                name: "",
                description: "",
                price: "",
                category: "",
                discount: "",
                availability_status: true,
                image_url: "",
            });

            if (newMenuId) {
                navigate(`/add-inventory/${newMenuId}`);
            }
        }
    });

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
        const { name, value, type } = e.target as HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement;
        const checked = type === "checkbox" ? (e.target as HTMLInputElement).checked : undefined;
        setFormData((prev) => ({
            ...prev,
            [name]: type === "checkbox" ? checked : value,
        }));
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        const updatedFormData = {
            ...formData,
            price: parseFloat(formData.price),
            discount: parseFloat(formData.discount),
        };

        console.log(updatedFormData);

        try {
            await addMenuItem({
                variables: {
                    input: updatedFormData,
                },
            });
        } catch (error) {
            console.error('Error during mutation:', error);
        } finally {
            console.log('Mutation is settled!');
        }
    };

    const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
        const file = e.target.files?.[0];
        if (!file) return;
    
        setFileUploading(true);
    
        try {
            const fileExt = file.name.split('.').pop();
            const fileName = `${userId}-${Date.now()}.${fileExt}`;
            const filePath = `menupic/${fileName}`;
    
            const { error: uploadError } = await supabase
                .storage
                .from('pictures')
                .upload(filePath, file);
    
            if (uploadError) {
                console.error('Error uploading file:', uploadError.message);
                return;
            }
    
            const { data: publicUrlData } = supabase
                .storage
                .from('pictures')
                .getPublicUrl(filePath);
    
            const publicUrl = publicUrlData.publicUrl;
    
            setFormData((prev) => ({
                ...prev,
                image_url: publicUrl,
            }));
        } catch (err) {
            console.error('Upload error:', err);
        } finally {
            setFileUploading(false);
        }
    };

    const HandleNavigateDashboard = () => {
        navigate("/dashboard");
    }

  return (
    <><Navbar />
    <form onSubmit={handleSubmit} className="max-w-150 mx-auto p-4 shadow-md rounded-xl bg-white space-y-4 text-black mt-4">
          <h2 className="text-xl font-bold">Add New Menu Item</h2>

          <div className="flex flex-col space-y-1">
              <label htmlFor="name" className="font-medium">Name</label>
              <input id="name" name="name" value={formData.name} onChange={handleChange} className="input text-black mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md" placeholder="e.g. Spaghetti" required />
          </div>

          <div className="flex flex-col space-y-1">
              <label htmlFor="description" className="font-medium">Description</label>
              <textarea id="description" name="description" value={formData.description} onChange={handleChange} className="input text-black mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md" placeholder="Enter description..." />
          </div>

          <div className="flex flex-col space-y-1">
              <label htmlFor="price" className="font-medium">Price (₱)</label>
              <input id="price" type="number" name="price" value={formData.price} onChange={handleChange} className="input text-black mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md" step="0.01" placeholder="e.g. 9.99" required />
          </div>

          <div className="flex flex-col space-y-1">
              <label htmlFor="category" className="font-medium">Category</label>
              <select id="category" name="category" value={formData.category} onChange={handleChange} className="input text-black mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md">
                  <option value="">Select a category</option>
                  <option value="Appetizer">Appetizer</option>
                  <option value="Main Course">Main Course</option>
                  <option value="Dessert">Dessert</option>
                  <option value="Beverage">Beverage</option>
                  <option value="Soup">Soup</option>
                  <option value="Salad">Salad</option>
                  <option value="Sides">Sides</option>
                  <option value="Vegan">Vegan</option>
                  <option value="Kids Meal">Kids Menu</option>
              </select>
          </div>

          <div className="flex flex-col space-y-1">
              <label htmlFor="discount" className="font-medium">Discount (%)</label>
              <input id="discount" type="number" name="discount" value={formData.discount} onChange={handleChange} className="input text-black mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md" step="0.01" placeholder="e.g. 10" required />
          </div>

          <div className="flex flex-col space-y-1">
              <label htmlFor="image_file" className="font-medium">Upload Image</label>
              <input id="image_file" title="image file" type="file" accept="image/*" onChange={handleFileChange} className="input text-black mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md" />
          </div>

          <div className="flex items-center space-x-2">
              <input id="availability_status" type="checkbox" name="availability_status" checked={formData.availability_status} onChange={handleChange} />
              <label htmlFor="availability_status" className="text-black">Available?</label>
          </div>

          <div className="flex space-x-2">
              <button type="submit" disabled={loading || fileUploading} className="btn btn-primary w-1/2 text-amber-50">
                  {loading || fileUploading ? "Loading..." : "Add Menu Item"}
              </button>
              <button className="btn btn-primary w-1/2 text-amber-50"
              onClick={HandleNavigateDashboard}>
                  Back
              </button>
          </div>

          {error && <p className="text-red-500 text-sm">Error: {error.message}</p>}
      </form></>
  );  
};

export default AddMenu;
