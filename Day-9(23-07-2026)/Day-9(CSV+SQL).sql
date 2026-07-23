use  golangpractice;
select * from evvehicles;
desc evvehicles;
SET SQL_SAFE_UPDATES = 0;

-- 1.Data Cleaning: Standardize manufacture_date into a single YYYY-MM-DD format (it currently has 3 different formats mixed in). Then find all vehicles manufactured after Jan 1, 2024.
UPDATE evvehicles
SET manufacture_date = DATE_FORMAT(STR_TO_DATE(manufacture_date,'%d/%m/%Y'),'%Y-%m-%d')
WHERE manufacture_date LIKE '__/__/____';

UPDATE evvehicles
SET manufacture_date = DATE_FORMAT(
    STR_TO_DATE(manufacture_date,'%m-%d-%Y'),
    '%Y-%m-%d')
WHERE manufacture_date LIKE '__-__-____';

UPDATE evvehicles
SET manufacture_date = DATE_FORMAT(STR_TO_DATE(manufacture_date,'%d-%m-%Y'),'%Y-%m-%d')
WHERE manufacture_date LIKE '__-__-____';

select * from evvehicles;
select * from evvehicles where manufacture_date > '2024-01-01';

-- 2.Grouped Aggregation: Calculate the average range_km per company_name, sorted descending. Which manufacturer has the best average range?
select company_name,avg(range_km) from evvehicles group by company_name order by avg(range_km) desc limit 1;

-- 3.Missing Data Handling: Identify all rows with missing values in warranty_years, odometer_km, or is_certified_pre_owned. Decide and justify a strategy (drop vs. impute) for each column, then apply it.
update evvehicles set odometer_km=(
select avg(odometer_km) from (select odometer_km from evvehicles where odometer_km is not null) as t)
where odometer_km is null;

select is_certified_pre_owned
 from evvehicles;

select is_certified_pre_owned,count(*) from evvehicles where is_certified_pre_owned is not null group by is_certified_pre_owned;

-- 4.Multi-condition Filter: Find all vehicles where battery_chemistry == "LFP" AND price_usd < 50000 AND region_sold is either "India" or "Asia Pacific" — these represent budget-friendly LFP EVs in key growth markets.
select * from evvehicles where battery_chemistry = 'LFP' AND price_usd < 50000 AND (region_sold ="India" or region_sold="Asia Pacific") ;

-- 5.Derived Metric: Create a new computed column price_per_kwh (price_usd / battery_capacity_kwh) and price_per_km_range (price_usd / range_km). Rank the top 5 most cost-efficient vehicles by price_per_km_range.
select model_name,price_usd / battery_capacity_kwh as price_per_kwh,price_usd / range_km as price_per_km_range from evvehicles order by price_per_km_range limit 5; 