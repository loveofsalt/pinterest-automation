# Pinterest Pin Creator - GitHub Actions Setup (Safe Batch Processing)

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
├── pins/               # Batch CSV files
│   ├── TEMPLATE_batch_YYYY_MM_DD.csv  # Template for new batches
│   ├── batch_2026_01_07.csv          # Today's batch
│   ├── batch_2026_01_14.csv          # Next week's batch
│   └── archive/                      # Processed batches
│       └── batch_2026_01_07_processed_20260107_143022.csv
└── sample_pins.csv     # Example CSV (for testing)
```

### 2. Set GitHub Secrets

Go to your repository → Settings → Secrets and Variables → Actions, and add:

- `PINTEREST_APP_ID` - Your Pinterest app ID
- `PINTEREST_APP_SECRET` - Your Pinterest app secret  
- `PINTEREST_REFRESH_TOKEN` - Your refresh token
- `PINTEREST_BOARD_ID` - Target Pinterest board ID

### 3. Safe Batch Workflow

#### Step 1: Create a New Batch
1. Copy `pins/TEMPLATE_batch_YYYY_MM_DD.csv` 
2. Rename to current date: `pins/batch_2026_01_07.csv`
3. Fill in your pin details
4. Commit and push

#### Step 2: Run Manual Processing
1. Go to Actions tab → "Pinterest Batch Pin Creator"
2. Click "Run workflow"
3. Enter: `pins/batch_2026_01_07.csv`
4. Choose archive option (recommended: true)
5. Click "Run workflow"

#### Step 3: Automatic Archival
- ✅ Successful processing moves CSV to `pins/archive/`
- 🏷️ Adds timestamp to filename
- 📝 Auto-commits archive with descriptive message
- 🔄 Ready for next batch

### 4. Why This Approach is Safest

- ✅ **No Duplicate Uploads**: Each CSV processes only once
- ✅ **Manual Control**: You decide exactly when pins are created  
- ✅ **Clear History**: Archived files show what was processed when
- ✅ **No Overwrites**: Fresh CSV for each batch prevents confusion
- ✅ **Rollback Friendly**: Archive contains exact pins that were uploaded

### 5. Best Practices

#### Naming Convention
```
pins/batch_2026_01_07.csv        # Today's batch
pins/batch_2026_01_14.csv        # Next week's batch
pins/holiday_batch_2026_12_25.csv # Special batches
```

#### Workflow
1. **Monday**: Create `batch_2026_01_13.csv` for next week
2. **Throughout week**: Add pins to next week's batch
3. **Friday**: Run workflow for current week's batch
4. **Archive**: Let GitHub automatically archive the processed file

#### File Organization
- **Active batches**: Keep in `pins/` directory
- **Processed batches**: Auto-moved to `pins/archive/`
- **Templates**: Use `TEMPLATE_` prefix for reusable templates