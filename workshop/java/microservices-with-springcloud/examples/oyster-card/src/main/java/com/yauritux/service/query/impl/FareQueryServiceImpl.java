package com.yauritux.service.query.impl;

import java.util.List;
import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.yauritux.model.constant.JourneyType;
import com.yauritux.model.entity.Fare;
import com.yauritux.repository.FareRepository;
import com.yauritux.service.query.FareQueryService;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
@Service
public class FareQueryServiceImpl implements FareQueryService {
	
	private FareRepository fareRepository;
	
	@Autowired
	public FareQueryServiceImpl(FareRepository fareRepository) {
		this.fareRepository = fareRepository;
	}

	@Override
	public Optional<List<Fare>> findByJourneyType(JourneyType journeyType) {
		return Optional.ofNullable(fareRepository.findByJourneyType(journeyType));
	}
}
