package asset

type Exporter interface {
	// ExportArea(ctx context.Context, model *ent.AreaReleaseAsset) error
}

// func (r *Repository) ExportArea(ctx context.Context, model *ent.AreaReleaseAsset) error {
// 	if model.Status != 1 {
// 		return nil
// 	}
// 	if model.FilePath == "" {
// 		return nil
// 	}
// 	filePath := filepath.Join(path.RootPath(), model.FilePath)
// 	f, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	reader := csv.NewReader(f)
// 	for {
// 		row, err := reader.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		rowLen := len(row)
// 		if rowLen == 0 {
// 			return errors.New("empty row")
// 		}
// 		var data = []string{}
// 		for _, value := range row {
// 			value = strings.TrimSpace(value)
// 			data = append(data, value)
// 		}
// 		id := data[0]
// 		if id == "" {
// 			return errors.New("empty row")
// 		}
// 		id = strings.ToLower(id)
// 		if id == "id" {
// 			return errors.New("empty row")
// 		}
// 		levelIdLen := 10
// 		// ID长度补全
// 		idUint, err := strconv.ParseUint(id, 10, 64)
// 		if err != nil {
// 			return err
// 		}
// 		id = strconv.FormatUint(idUint, 10)
// 		idLen := len(id)
// 		if idLen > levelIdLen {
// 			id = id[0:levelIdLen]
// 		} else if idLen < levelIdLen {
// 			id = id + strings.Repeat("0", levelIdLen-idLen)
// 		}
// 		// PID长度补全
// 		pidUint, err := strconv.ParseUint(data[1], 10, 64)
// 		if err != nil {
// 			return err
// 		}
// 		pid := strconv.FormatUint(pidUint, 10)
// 		pidLen := len(pid)
// 		if pidLen > levelIdLen {
// 			pid = pid[0:levelIdLen]
// 		} else if pidLen < levelIdLen {
// 			pid = pid + strings.Repeat("0", levelIdLen-pidLen)
// 		}
// 		// 深度和级别
// 		level, err := strconv.ParseInt(data[2], 10, 64)
// 		if err != nil {
// 			return err
// 		}
// 		// 名称
// 		name := data[3]
// 		// 拼音
// 		pinyin := data[5]
// 		if pinyin == "" || pinyin[0:1] == "-" {
// 			pinyin = ""
// 		} else {
// 			pinyin = cases.Title(language.Make(pinyin)).String(pinyin)
// 		}
// 		// 原始ID
// 		extId := data[6]
// 		// 原始名称
// 		extName := data[7]
// 		// 开始导入数据
// 		ok, err := tx.Area.Query().
// 		Where(area.RegionIDEQ(r.ID), area.LevelEQ(uint8(r.Level))).
// 		Exist(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	if ok {
// 		_, err = tx.Area.Update().
// 			Where(area.RegionIDEQ(r.ID), area.LevelEQ(uint8(r.Level))).
// 			SetRegionID(r.ID).
// 			SetTitle(r.Name).
// 			SetPinyin(r.Pinyin).
// 			SetUcfirst(r.Ucfirst()).
// 			SetLevel(uint8(r.Level)).
// 			Save(ctx)
// 	} else {
// 		_, err = tx.Area.Create().
// 			SetRegionID(r.ID).
// 			SetTitle(r.Name).
// 			SetPinyin(r.Pinyin).
// 			SetUcfirst(r.Ucfirst()).
// 			SetLevel(uint8(r.Level)).
// 			Save(ctx)
// 	}
// 	if err != nil {
// 		return err
// 	}
// 	}
// 	return nil
// }
