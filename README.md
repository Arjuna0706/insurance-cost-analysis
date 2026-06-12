# 💰 Medical Insurance Cost Analysis & Prediction

> Predicting individual medical insurance charges from demographic and health attributes using regression modeling — completed as a capstone project for **IBM's *Data Analysis with Python*** course on Coursera.

![Python](https://img.shields.io/badge/Python-3.x-3776AB?style=flat&logo=python&logoColor=white)
![pandas](https://img.shields.io/badge/pandas-Data%20Wrangling-150458?style=flat&logo=pandas&logoColor=white)
![scikit-learn](https://img.shields.io/badge/scikit--learn-ML-F7931E?style=flat&logo=scikit-learn&logoColor=white)
![Jupyter](https://img.shields.io/badge/Jupyter-Notebook-F37626?style=flat&logo=jupyter&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green.svg)

---

## 📌 Overview

This project builds a machine learning model to estimate the **annual medical insurance charges** billed to an individual, based on personal attributes such as age, BMI, smoking status, and region. Using a dataset of **2,772 policyholders**, the analysis follows the full data-science lifecycle — from data preparation and exploratory analysis, through model development, to evaluation and refinement.

The central question: **which personal and health factors drive insurance costs the most, and how accurately can we predict an individual's charges from them?**

This kind of model has real-world value for insurers and healthcare providers in **risk assessment** and **data-driven premium pricing**.

---

## 🎯 Objectives

- Prepare and encode a real-world insurance dataset for modeling.
- Explore how individual features — especially BMI and smoking — relate to charges.
- Build regression models of increasing complexity and benchmark their accuracy.
- Apply polynomial features and Ridge regularization to capture feature interactions and improve generalization.

---

## 📊 Dataset

The dataset contains **2,772 records** across **7 columns**, with each row representing one policyholder. The data is clean, with **no missing values**.

| Feature | Description |
|---|---|
| `charges` | **Target variable** — annual medical insurance cost billed (USD) |
| `age` | Age of the primary beneficiary |
| `sex` | Gender (`female` / `male`) |
| `bmi` | Body Mass Index — a measure of body fat based on height and weight |
| `children` | Number of dependents covered |
| `smoker` | Smoking status (`yes` / `no`) |
| `region` | Residential region in the US (NE, NW, SE, SW) |

**Charges range** from \$1,122 to \$63,770, with a mean of **~\$13,261** — a wide and right-skewed distribution that the modeling needs to account for.

---

## 🛠️ Tech Stack

| Category | Tools |
|---|---|
| **Language** | Python 3 |
| **Data handling** | pandas, NumPy |
| **Visualization** | Matplotlib, Seaborn |
| **Modeling** | scikit-learn (`LinearRegression`, `Ridge`, `PolynomialFeatures`, `Pipeline`, `StandardScaler`) |
| **Data source** | `kagglehub` (programmatic Kaggle dataset download) |
| **Environment** | Jupyter Notebook |

---

## 🔍 Workflow & Methodology

### 1. Data Preparation
- Loaded the dataset directly from Kaggle via `kagglehub`.
- Replaced placeholder `?` values with `NaN`.
- **Encoded categorical variables** into numeric form: `sex` (female=0, male=1), `smoker` (no=0, yes=1), and `region` (0–3).
- Cast `bmi` to a proper float type.

### 2. Exploratory Data Analysis (EDA)
- **BMI vs Charges (regplot):** revealed a positive but loosely-dispersed relationship — higher BMI tends to mean higher charges, but BMI alone doesn't tell the full story.
- **Smoker vs Charges (boxplot):** exposed a dramatic gap — the *lower* bound of the smoker group nearly aligns with the *upper outliers* of non-smokers.
- **Correlation matrix:** quantified each feature's relationship with charges.

### 3. Model Development & Refinement
Models were built progressively to show the impact of feature richness, non-linearity, and regularization:

| Step | Model | Features | R² Score |
|---|---|---|:---:|
| 1 | Simple Linear Regression | `smoker` only | 0.622 |
| 2 | Multiple Linear Regression | All 6 features | 0.751 |
| 3 | Polynomial Pipeline (scale → poly° 2 → linear) | All features (train) | **0.846** |
| 4 | Ridge Regression (α=0.1) | All features (test set) | 0.678 |
| 5 | Ridge + 2nd-degree Polynomial (test set) | All features | **0.786** |

> Steps 1–3 are evaluated in-sample; steps 4–5 use a proper **80/20 train-test split** for an honest estimate of out-of-sample performance.

---

## 💡 Key Insights

- **Smoking is the single most powerful predictor.** It has a correlation of **0.79** with charges — far above any other feature. On average, smokers are billed **~\$32,223** versus **~\$8,418** for non-smokers — almost **4× higher**.
- **Age and BMI matter, but secondarily.** Age (corr ≈ 0.30) and BMI (corr ≈ 0.20) contribute meaningfully, but neither comes close to the impact of smoking status.
- **Region and sex are nearly irrelevant.** With correlations near zero (0.01 and 0.06), geographic region and gender add almost no predictive power on their own.
- **Feature interactions are the key to accuracy.** Moving from a single feature to all six lifted R² from **0.62 → 0.75**. Adding 2nd-degree polynomial features pushed it further to **0.85** in-sample — capturing combined effects like *the compounding cost of being both a smoker and high-BMI*.
- **The model generalizes well.** On unseen test data, the polynomial + Ridge model holds at **R² ≈ 0.79**, confirming it learns real patterns rather than memorizing the training set.

---

## 📈 Results Summary

The best-performing pipeline — **StandardScaler → PolynomialFeatures (degree 2) → Linear Regression** — explains roughly **85% of the variance** in insurance charges. On held-out test data, the regularized polynomial model maintains a strong **R² ≈ 0.79**, making it a reliable baseline for automated insurance cost estimation.

---

## 🚀 Getting Started

### Prerequisites
```bash
pip install pandas numpy matplotlib seaborn scikit-learn kagglehub jupyter
```

### Run the notebook
```bash
git clone https://github.com/<your-username>/<your-repo>.git
cd <your-repo>
jupyter notebook Project_Insurance_Cost_Analysis1.ipynb
```

> **Note:** The notebook downloads the dataset automatically via `kagglehub`. A local copy (`Medical_insurance.csv`) is also included in this repository if you prefer to load it directly.

---

## 📁 Repository Structure

```
.
├── Project_Insurance_Cost_Analysis1.ipynb   # Main analysis notebook
├── Medical_insurance.csv                     # Insurance dataset (2,772 records)
└── README.md                                 # Project documentation
```

---

## 🔮 Future Improvements

- Apply a **log-transformation** to `charges` to address its right-skew and stabilize residuals.
- Engineer an explicit **smoker × BMI interaction term** to directly model the strongest combined effect.
- Test ensemble methods (**Random Forest**, **Gradient Boosting / XGBoost**) for likely accuracy gains.
- Add cross-validated hyperparameter tuning (`GridSearchCV`) for the Ridge α and polynomial degree.
- Report additional metrics (**RMSE, MAE**) alongside R² for a fuller picture of error in dollar terms.

---

## 🙏 Acknowledgements

- **Dataset:** [Medical Insurance Price Prediction](https://www.kaggle.com/datasets/harishkumardatalab/medical-insurance-price-prediction) by Harish Kumar (Kaggle).
- **Course:** *Data Analysis with Python* — IBM, via Coursera.

---

## 📄 License

This project is released under the **MIT License** — feel free to use, modify, and learn from it.
