package com.yauritux.service.query.impl;

import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.yauritux.model.entity.Station;
import com.yauritux.repository.StationRepository;
import com.yauritux.service.query.StationQueryService;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@Service
public class StationQueryServiceImpl implements StationQueryService {
	
	private StationRepository stationRepository;
	
	@Autowired
	StationQueryServiceImpl(StationRepository stationRepository) {
		this.stationRepository = stationRepository;
	}

	@Override
	public Optional<Station> findByName(String name) {
		return Optional.ofNullable(stationRepository.findByName(name));
	}

}
