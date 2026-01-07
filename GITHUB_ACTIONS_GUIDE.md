# Pinterest Pin Creator - GitHub Actions Setup

## How to Use with GitHub Actions

### 1. Repository Structure
```
your-repo/
├── .github/workflows/pinterest-batch.yml
├── main.go
├── check_images.go
├── images/              # Your image files
│   ├── recipe1.jpg
│   ├── recipe2.jpg
│   └── recipe3.jpg
├── pins/               # CSV files for batch processing
│   ├── recipes_batch.csv
│   └── weekly_pins.csv
└── sample_pins.csv     # Example CSV
```

### 2. Set GitHub Secrets

Go to your repository → Settings → Secrets and Variables → Actions, and add:

- `PINTEREST_APP_ID` - Your Pinterest app ID
- `PINTEREST_APP_SECRET` - Your Pinterest app secret  
- `PINTEREST_REFRESH_TOKEN` - Your refresh token
- `PINTEREST_BOARD_ID` - Target Pinterest board ID

### 3. Usage Options

#### Option A: Manual Trigger
1. Go to Actions tab in your GitHub repository
2. Select "Pinterest Batch Pin Creator" workflow
3. Click "Run workflow"
4. Specify CSV file path (e.g., `pins/recipes_batch.csv`)
5. Click "Run workflow"

#### Option B: Automatic on CSV Changes
1. Add/update any CSV file in the `pins/` directory
2. Commit and push to main branch
3. Workflow automatically runs

#### Option C: Scheduled Batches
Add to your workflow file:
```yaml
on:
  schedule:
    - cron: '0 9 * * 1'  # Every Monday at 9 AM UTC
```

### 4. Workflow Features

- ✅ **Validation**: Checks that all image files exist before processing
- 📊 **Progress**: Shows detailed progress and results
- 🛡️ **Error Handling**: Continues processing even if individual pins fail
- 📋 **CSV Preview**: Shows CSV contents in the workflow log
- 🔄 **Flexible Triggering**: Manual, automatic, or scheduled

### 5. Example Workflow Run

```bash
# Workflow validates your CSV
✅ Found CSV file: pins/recipes_batch.csv
📋 CSV Contents:
file_path,title,description,link,alt_text,section_id,note
images/recipe1.jpg,Salt-Baked Fish,Delicious recipe,,Fresh fish,,Amazing!

# Checks all images exist
✅ Found: images/recipe1.jpg
✅ Found: images/recipe2.jpg
🎉 All image files found!

# Processes pins
🔄 Authenticating...
📂 Processing CSV file: pins/recipes_batch.csv
📊 Found 3 pins to process
🔄 Processing pin 1/3: Salt-Baked Fish
✅ Pin 1/3 created successfully: recipe1.jpg
✅ Batch processing complete! Success: 3, Failed: 0
```

### 6. Best Practices

1. **Organize by Campaign**: Create separate CSV files for different pin campaigns
2. **Test First**: Use `sample_pins.csv` to test your setup
3. **Image Paths**: Use relative paths from repo root (e.g., `images/photo.jpg`)
4. **Default Links**: Empty link fields will automatically use `https://www.loveofsalt.com`
5. **Commit Images**: Ensure all image files are committed to the repository

### 7. Troubleshooting

- **Missing Images**: The workflow will fail if any image in CSV doesn't exist
- **Pinterest API Limits**: Consider adding delays between pins for large batches
- **CSV Format**: Ensure proper CSV format (commas, quotes for fields with commas)
- **File Paths**: Use forward slashes `/` even on Windows