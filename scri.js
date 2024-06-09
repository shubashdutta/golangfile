document
  .getElementById("createFolderBtn")
  .addEventListener("click", async function () {
    const folderName = document.getElementById("folderName").value.trim();

    if (folderName === "") {
      document.getElementById("message").textContent =
        "Please enter a folder name.";
      return;
    }

    try {
      const response = await fetch(
        "https://golangfile-2.onrender.com/api/create-folder",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ folderName }),
        }
      );

      if (response.ok) {
        const data = await response.json();
        location.reload();
        document.getElementById("message").textContent = data.message;
      } else {
        const errorData = await response.json();
        throw new Error(errorData.error);
      }
    } catch (error) {
      console.error("Error creating folder:", error);
      document.getElementById("message").textContent =
        "Error creating folder. Please try again.";
    }
  });
